package repository

import (
	"os"
	"path/filepath"

	"plots/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Repositories 仓库集合
type Repositories struct {
	User      *UserRepository
	Article   *ArticleRepository
	Portfolio *PortfolioRepository
}

// NewRepositories 创建仓库集合
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:      NewUserRepository(db),
		Article:   NewArticleRepository(db),
		Portfolio: NewPortfolioRepository(db),
	}
}

// InitDB 初始化数据库
func InitDB(dbPath string) (*gorm.DB, error) {
	// 确保数据库目录存在
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// 打开数据库连接
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// 获取底层 sql.DB 并配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SQLite 配置
	sqlDB.SetMaxOpenConns(1) // SQLite 不支持多连接写入
	sqlDB.SetMaxIdleConns(1)

	return db, nil
}

// ========================================
// UserRepository 用户仓库
// ========================================

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	return count, err
}

// ========================================
// ArticleRepository 文章仓库
// ========================================

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Create(article *model.Article) error {
	return r.db.Create(article).Error
}

func (r *ArticleRepository) GetByID(id uint) (*model.Article, error) {
	var article model.Article
	err := r.db.Preload("Author").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) Update(article *model.Article) error {
	return r.db.Save(article).Error
}

func (r *ArticleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}

func (r *ArticleRepository) List(query model.PaginationQuery, publishedOnly bool) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	db := r.db.Model(&model.Article{})

	// 只显示已发布的
	if publishedOnly {
		db = db.Where("published = ?", true)
	}

	// 分类筛选
	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}

	// 搜索
	if query.Search != "" {
		search := "%" + query.Search + "%"
		db = db.Where("title LIKE ? OR summary LIKE ?", search, search)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (query.Page - 1) * query.PageSize
	err := db.Preload("Author").
		Order("created_at DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&articles).Error

	return articles, total, err
}

func (r *ArticleRepository) IncrementViews(id uint) error {
	return r.db.Model(&model.Article{}).Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}

// ========================================
// PortfolioRepository 作品集仓库
// ========================================

type PortfolioRepository struct {
	db *gorm.DB
}

func NewPortfolioRepository(db *gorm.DB) *PortfolioRepository {
	return &PortfolioRepository{db: db}
}

func (r *PortfolioRepository) Create(portfolio *model.Portfolio) error {
	return r.db.Create(portfolio).Error
}

func (r *PortfolioRepository) GetByID(id uint) (*model.Portfolio, error) {
	var portfolio model.Portfolio
	err := r.db.First(&portfolio, id).Error
	if err != nil {
		return nil, err
	}
	return &portfolio, nil
}

func (r *PortfolioRepository) Update(portfolio *model.Portfolio) error {
	return r.db.Save(portfolio).Error
}

func (r *PortfolioRepository) Delete(id uint) error {
	return r.db.Delete(&model.Portfolio{}, id).Error
}

func (r *PortfolioRepository) List(publishedOnly bool) ([]model.Portfolio, error) {
	var portfolios []model.Portfolio

	db := r.db.Model(&model.Portfolio{})

	if publishedOnly {
		db = db.Where("published = ?", true)
	}

	err := db.Order("sort_order ASC, created_at DESC").Find(&portfolios).Error
	return portfolios, err
}
