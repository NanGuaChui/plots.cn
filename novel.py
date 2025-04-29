# 导入tkinter库
import threading
import tkinter as tk
import pystray
from pystray import MenuItem as item
from PIL import Image, ImageDraw
import json
from tkinter import colorchooser

# 创建托盘图标图像
icon_image = Image.new("RGB", (64, 64), (255, 255, 255))
draw = ImageDraw.Draw(icon_image)
draw.rectangle((16, 16, 48, 48), fill="black")


def create_tray_icon():
    def show_hide_window():
        if root.state() == "withdrawn":
            root.deiconify()
        else:
            root.withdraw()

    def quit_app():
        tray_icon.stop()
        root.destroy()

    # 创建托盘菜单
    menu = (
        item("显示/隐藏", lambda: show_hide_window()),
        item("设置", lambda: open_settings()),
        item("退出", lambda: quit_app()),
    )

    tray_icon = pystray.Icon("SimpleApp", icon_image, "简单计算机", menu)
    return tray_icon


def center_window(window):
    """将窗口移动到屏幕中央"""
    window.update_idletasks()
    screen_width = window.winfo_screenwidth()
    screen_height = window.winfo_screenheight()
    window_width = window.winfo_width()
    window_height = window.winfo_height()
    x = (screen_width - window_width) // 2
    y = (screen_height - window_height) // 2
    window.geometry(f"{window_width}x{window_height}+{x}+{y}")


def make_window_draggable(window):
    """使窗口支持任意位置拖动"""
    def start_drag(event):
        window._drag_data = {"x": event.x, "y": event.y}

    def do_drag(event):
        x = window.winfo_x() + (event.x - window._drag_data["x"])
        y = window.winfo_y() + (event.y - window._drag_data["y"])
        window.geometry(f"+{x}+{y}")

    window.bind("<Button-1>", start_drag)
    window.bind("<B1-Motion>", do_drag)


# 创建主窗口
root = tk.Tk()
root.title("云黑课堂例子——简单计算机")


# 设置窗口无边框和无顶部菜单栏
root.overrideredirect(True)

# 设置窗口始终显示在最前
root.attributes('-topmost', True)


# 居中窗口
center_window(root)

# 启用窗口拖动功能
make_window_draggable(root)

# 创建托盘图标线程
tray_icon = create_tray_icon()

# 启动托盘图标
tray_thread = threading.Thread(target=tray_icon.run, daemon=True)
tray_thread.start()


def load_novel(file_path):
    """加载小说内容并按换行符和最大字数分页"""
    max_chars = 100  # 每页最大字数
    pages = []
    with open(file_path, "r", encoding="utf-8") as file:
        lines = file.readlines()
        for line in lines:
            line = line.strip()
            while len(line) > max_chars:
                pages.append(line[:max_chars])
                line = line[max_chars:]
            if line:  # 添加剩余部分
                pages.append(line)
    return pages


def display_page(page_index):
    """显示指定页内容并动态调整窗口高度和宽度"""
    if 0 <= page_index < len(pages):
        page_content = f"{pages[page_index]}    {page_index + 1}/{len(pages)}"
        label.config(text=page_content)
        label.update_idletasks()

        # 动态调整高度
        content_height = label.winfo_reqheight()
        new_height = min(
            content_height, root.winfo_screenheight() - 100)  # 限制最大高度

        # 动态调整宽度
        content_width = label.winfo_reqwidth()
        new_width = content_width  # 动态调整宽度

        root.geometry(f"{new_width}x{new_height}")


def next_page(event=None):
    """显示下一页"""
    global current_page
    if current_page < len(pages) - 1:
        current_page += 1
        display_page(current_page)


def previous_page(event=None):
    """显示上一页"""
    global current_page
    if current_page > 0:
        current_page -= 1
        display_page(current_page)


# 加载小说内容
novel_file_path = "novel.txt"
pages = load_novel(novel_file_path)
current_page = 0

# 创建标签用于显示内容
label = tk.Label(root, font=("Arial", 12), anchor="nw",
                 justify="left", padx=0, pady=0)
label.pack(expand=True, fill="both")

# 显示第一页内容
display_page(current_page)

# 绑定方向键事件
root.bind("<Right>", next_page)
root.bind("<Left>", previous_page)

# 加载配置文件
config_file = "config.json"
try:
    with open(config_file, "r", encoding="utf-8") as f:
        config = json.load(f)
except FileNotFoundError:
    config = {}

text_color = config.get("text_color", "#000000")
bg_color = config.get("bg_color", "#FFFFFF")


def save_config():
    """保存配置到config.json"""
    config["text_color"] = text_color
    config["bg_color"] = bg_color
    with open(config_file, "w", encoding="utf-8") as f:
        json.dump(config, f, ensure_ascii=False, indent=4)


def open_settings():
    """打开设置对话框"""
    def choose_text_color(event=None):
        global text_color
        color = colorchooser.askcolor(
            title="选择文字颜色", initialcolor=text_color[:7])[1]  # 去掉透明度部分
        if color:
            text_color = color + text_color[7:]  # 保留透明度部分
            label.config(fg=color)  # 使用无透明度的颜色
            text_color_block.config(bg=color)  # 更新文字颜色色块
            save_config()

    def choose_bg_color(event=None):
        global bg_color
        color = colorchooser.askcolor(
            title="选择背景颜色", initialcolor=bg_color[:7])[1]  # 去掉透明度部分
        if color:
            bg_color = color + bg_color[7:]  # 保留透明度部分
            label.config(bg=color)  # 使用无透明度的颜色
            root.config(bg=color)  # 使用无透明度的颜色
            bg_color_block.config(bg=color)  # 更新背景颜色色块
            save_config()

    def set_text_opacity(value):
        global text_color
        r, g, b = root.winfo_rgb(text_color[:7])[:3]  # 去掉透明度部分
        text_color = f"#{r >> 8:02x}{g >> 8:02x}{b >> 8:02x}{int(float(value) * 255):02x}"
        label.config(fg=text_color[:7])  # 使用无透明度的颜色
        save_config()

    def set_bg_opacity(value):
        global bg_color
        r, g, b = root.winfo_rgb(bg_color[:7])[:3]  # 去掉透明度部分
        bg_color = f"#{r >> 8:02x}{g >> 8:02x}{b >> 8:02x}{int(float(value) * 255):02x}"
        label.config(bg=bg_color[:7])  # 使用无透明度的颜色
        root.config(bg=bg_color[:7])  # 使用无透明度的颜色
        save_config()

    settings_window = tk.Toplevel(root)
    settings_window.title("设置")
    settings_window.geometry("300x200")
    settings_window.resizable(False, False)
    settings_window.attributes('-topmost', True)

    tk.Label(settings_window, text="文字颜色:").grid(
        row=0, column=0, padx=10, pady=5)
    text_color_block = tk.Label(
        settings_window, bg=text_color[:7], width=10, height=1)  # 去掉透明度部分
    text_color_block.grid(row=0, column=1, padx=10, pady=5)
    text_color_block.bind("<Button-1>", choose_text_color)  # 点击色块打开颜色选择器

    tk.Label(settings_window, text="背景颜色:").grid(
        row=1, column=0, padx=10, pady=5)
    bg_color_block = tk.Label(
        settings_window, bg=bg_color[:7], width=10, height=1)  # 去掉透明度部分
    bg_color_block.grid(row=1, column=1, padx=10, pady=5)
    bg_color_block.bind("<Button-1>", choose_bg_color)  # 点击色块打开颜色选择器

    tk.Label(settings_window, text="文字透明度:").grid(
        row=2, column=0, padx=10, pady=5)
    tk.Scale(settings_window, from_=0, to=1, resolution=0.01, orient="horizontal",
             command=set_text_opacity).grid(row=2, column=1, columnspan=2, padx=10, pady=5)

    tk.Label(settings_window, text="背景透明度:").grid(
        row=3, column=0, padx=10, pady=5)
    tk.Scale(settings_window, from_=0, to=1, resolution=0.01, orient="horizontal",
             command=set_bg_opacity).grid(row=3, column=1, columnspan=2, padx=10, pady=5)


# 更新托盘菜单
menu = (
    item("显示/隐藏", lambda: show_hide_window()),
    item("设置", lambda: open_settings()),
    item("退出", lambda: quit_app()),
)
tray_icon = pystray.Icon("SimpleApp", icon_image, "简单计算机", menu)


def create_context_menu():
    """创建右键菜单"""
    context_menu = tk.Menu(root, tearoff=0)
    context_menu.add_command(label="设置", command=open_settings)
    context_menu.add_command(label="退出", command=lambda: root.destroy())

    def show_context_menu(event):
        context_menu.post(event.x_root, event.y_root)

    root.bind("<Button-3>", show_context_menu)


# 创建右键菜单
create_context_menu()


# 应用初始配置
label.config(fg=text_color, bg=bg_color)
root.config(bg=bg_color)

# 启动主循环
root.mainloop()
