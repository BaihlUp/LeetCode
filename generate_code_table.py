import os
import re

def natural_sort_key(s):
    """
    定义自然排序的 key 函数
    """
    return [int(text) if text.isdigit() else text.lower() for text in re.split(r'(\d+)', s)]

def generate_code_table(directory):
    table = "| 文件名 | 描述 |\n|--------------|---------------------|\n"
    files = sorted(os.listdir(directory), key=natural_sort_key) # 按编号排序
    
    for filename in files:
        if filename.endswith(".go") or filename.endswith(".js"):
            filepath = os.path.join(directory, filename)
            filename = filename.rstrip(".go")
            table += f"| {filename} | [Go]({filepath}) |\n"
    
    return table

if __name__ == "__main__":
    code_directory = "Go/"
    readme_content = f"# 项目名称\n\nLeetCode C++/Go/Python 代码\n\n## 代码文件\n\n{generate_code_table(code_directory)}"
    
    with open("README.md", "w") as f:
        f.write(readme_content)
