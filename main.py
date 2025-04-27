import os
import json
import tkinter as tk
from tkinter import messagebox, filedialog
from tkinter.colorchooser import askcolor  # 确保导入 askcolor
from pystray import Icon, MenuItem, Menu
from PIL import Image, ImageTk
import threading

# 全局变量
current_page = 0
page_content = []

def read_config():
    """从config.json文件读取目录路径、文字颜色和页码信息"""
    config_path = "config.json"
    if not os.path.exists(config_path):
        return {"directory": None, "text_color": "black", "current_page": 0}
    try:
        with open(config_path, "r", encoding="utf-8") as f:
            config = json.load(f)
            return {
                "directory": config.get("directory"),
                "text_color": config.get("text_color", "black"),
                "current_page": config.get("current_page", 0)
            }
    except json.JSONDecodeError:
        return {"directory": None, "text_color": "black", "current_page": 0}

def save_config(directory=None, text_color=None, current_page=None):
    """保存目录路径、文字颜色和页码信息到config.json"""
    config_path = "config.json"
    config = read_config()
    if directory is not None:
        config["directory"] = directory
    if text_color is not None:
        config["text_color"] = text_color
    if current_page is not None:
        config["current_page"] = current_page
    with open(config_path, "w", encoding="utf-8") as f:
        json.dump(config, f, ensure_ascii=False, indent=4)

def read_txt_file(file_path):
    """读取txt文件内容"""
    if not os.path.exists(file_path):
        return "未找到对应文件"
    with open(file_path, "r", encoding="utf-8") as f:
        content = f.read().strip()
        return content if content else "未找到对应文件"

def paginate_content(content, width, height):
    """根据宽高分页内容"""
    lines = content.splitlines()
    page_lines = max(1, height // 20)  # 确保 page_lines 至少为 1
    return [lines[i:i + page_lines] for i in range(0, len(lines), page_lines)]

def update_content_label(text):
    """更新内容显示区域"""
    content_label.config(text=text)

def update_page_info():
    """更新页码信息"""
    page_info_label.config(text=f"{current_page + 1}/{len(page_content)}" if page_content else "0/0")

def display_page():
    """显示当前页内容"""
    global page_content, current_page
    if page_content:
        update_content_label("\n".join(page_content[current_page]))
    else:
        update_content_label("没有内容可显示")
    update_page_info()  # 更新页码信息

def display_content():
    """获取目录并展示内容"""
    global page_content, current_page
    config = read_config()
    directory = config["directory"]
    text_color = config["text_color"]
    current_page = config["current_page"]  # 恢复页码
    update_text_color(text_color)  # 更新文字颜色
    if not directory:
        update_content_label("配置文件中未指定目录")
        return
    content = read_txt_file(directory)
    if content == "未找到对应文件":
        update_content_label(content)
        return
    page_content = paginate_content(content, content_label.winfo_width(), content_label.winfo_height())
    if current_page >= len(page_content):  # 防止页码超出范围
        current_page = 0
    display_page()

def change_page(step):
    """切换页码"""
    global current_page, page_content
    new_page = current_page + step
    if 0 <= new_page < len(page_content):
        current_page = new_page
        save_config(current_page=current_page)  # 保存当前页码
        display_page()

def on_exit(icon, item):
    """托盘图标退出时关闭程序"""
    icon.stop()
    root.destroy()  # 销毁主窗口，确保程序完全退出

def select_file(icon, item):
    """选择txt文件并保存路径到config.json"""
    file_path = filedialog.askopenfilename(
        title="选择TXT文件",
        filetypes=[("文本文件", "*.txt")],
        initialdir=os.getcwd()
    )
    if file_path:
        save_config(directory=file_path)
        display_content()

def hide_window(event=None):
    """隐藏窗口"""
    root.withdraw()

def show_window(icon, item):
    """显示窗口"""
    root.deiconify()

def set_text_color(icon, item):
    """设置文字颜色并保存到config.json"""
    color = askcolor(title="选择文字颜色")[1]  # 返回颜色代码
    if color:
        save_config(text_color=color)
        update_text_color(color)

def update_text_color(color):
    """更新内容显示区域和页码信息的文字颜色"""
    content_label.config(fg=color)  # 更新内容文字颜色
    page_info_label.config(fg=color)  # 更新页码文字颜色

# 修改托盘图标菜单，添加显示窗口功能
def setup_tray_icon():
    """设置托盘图标"""
    icon_path = "favicon.ico"
    if not os.path.exists(icon_path):
        messagebox.showerror("错误", "未找到托盘图标文件 favicon.ico")
        return None
    image = Image.open(icon_path)
    menu = Menu(
        MenuItem("显示窗口", lambda: show_window(tray_icon, None)),
        MenuItem("选择文件", lambda: select_file(tray_icon, None)),
        MenuItem("设置文字颜色", lambda: set_text_color(tray_icon, None)),  # 新增选项
        MenuItem("退出", lambda: on_exit(tray_icon, None))
    )
    return Icon("novel-reader", image, "novel-reader", menu)

def start_move(event):
    """记录鼠标按下时的位置"""
    root.x = event.x
    root.y = event.y

def do_move(event):
    """根据鼠标移动更新窗口位置"""
    new_x = root.winfo_x() + (event.x - root.x)
    new_y = root.winfo_y() + (event.y - root.y)
    root.geometry(f"+{new_x}+{new_y}")

def change_cursor_to_move(event):
    """将鼠标指针样式改为 move"""
    root.config(cursor="fleur")

def reset_cursor(event):
    """重置鼠标指针样式"""
    root.config(cursor="")

# 创建主窗口
root = tk.Tk()
root.title("novel-reader")

# 设置窗口启动时居中显示
def center_window(window):
    """将窗口移动到屏幕中央"""
    window.update_idletasks()
    screen_width = window.winfo_screenwidth()
    screen_height = window.winfo_screenheight()
    window_width = window.winfo_width()
    window_height = window.winfo_height()
    x = (screen_width - window_width) // 2
    y = (screen_height - window_height) // 2
    window.geometry(f"+{x}+{y}")

# 绑定鼠标事件以实现拖动功能
root.bind("<Button-1>", start_move)  # 鼠标左键按下
root.bind("<B1-Motion>", do_move)   # 鼠标左键拖动
root.bind("<Enter>", change_cursor_to_move)  # 鼠标进入窗口时
root.bind("<Leave>", reset_cursor)  # 鼠标离开窗口时

# 设置窗口无边框和无顶部菜单栏
root.overrideredirect(True)

# 设置窗口背景色透明
root.attributes('-transparentcolor', root['bg'])

# 设置窗口透明度 (0.0 完全透明, 1.0 不透明)
root.attributes('-alpha', 1.0)

# 设置窗口始终显示在最前
root.attributes('-topmost', True)

# 绑定窗口大小变化事件和快捷键
root.bind("<Left>", lambda event: change_page(-1))  # 左箭头切换到上一页
root.bind("<Right>", lambda event: change_page(1))  # 右箭头切换到下一页
# 绑定 Esc 键隐藏窗口
root.bind("<Escape>", hide_window)

# 创建内容和页码的容器
content_frame = tk.Frame(root)
content_frame.pack(fill=tk.BOTH, expand=True)

# 创建内容显示区域
content_label = tk.Label(content_frame, wraplength=root.winfo_screenwidth(), justify="left", anchor="nw")
content_label.pack(side=tk.LEFT, fill=tk.BOTH, expand=True, padx=10, pady=10)

# 创建页码信息显示区域
page_info_label = tk.Label(content_frame, anchor="e", justify="right", width=10)
page_info_label.pack(side=tk.RIGHT, padx=10, pady=10)

# 初始化显示内容
root.update_idletasks()
center_window(root)  # 调用居中函数
display_content()

# 创建托盘图标
tray_icon = setup_tray_icon()
if tray_icon:
    threading.Thread(target=tray_icon.run, daemon=True).start()

# 运行应用程序
root.mainloop()
if tray_icon:
    tray_icon.stop()
