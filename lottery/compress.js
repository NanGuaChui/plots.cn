import sharp from 'sharp';
import fs from 'fs';
import path from 'path';

// 源目录和目标目录
const sourceDir = 'C:\\Users\\sgao\\Documents\\workspace\\its avatar';
const targetDir = 'C:\\Users\\sgao\\Develop\\plots.cn\\src\\assets\\img\\avatar';

// 确保目标目录存在
if (!fs.existsSync(targetDir)) {
  fs.mkdirSync(targetDir);
}

// 读取源目录中的所有文件
fs.readdir(sourceDir, (err, files) => {
  if (err) {
    console.error('无法读取源目录:', err);
    return;
  }

  files.forEach(file => {
    const filePath = path.join(sourceDir, file);
    const extname = path.extname(file);
    const basename = path.basename(file, extname).replace(/[\s,]/g, '');
    const targetPath = path.join(targetDir, `${basename}${extname.toLowerCase()}`);

    // 使用 sharp 调整图片大小
    sharp(filePath)
      .resize(100, 100)
      .toFile(targetPath, (err, info) => {
        if (err) {
          console.error('无法调整图片大小:', err);
        } else {
          console.log(`已调整大小并保存: ${targetPath}`);
        }
      });
  });
});
