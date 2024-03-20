import os

def replace_text_in_file(file_path, search_text, replace_text):
    with open(file_path, 'r') as file:
        file_contents = file.read()

    file_contents = file_contents.replace(search_text, replace_text)

    with open(file_path, 'w') as file:
        file.write(file_contents)

def traverse_directory(directory):
    for dirpath, dirnames, filenames in os.walk(directory):
        for filename in filenames:
            file_path = os.path.join(dirpath, filename)
            replace_text_in_file(file_path, 'Akssgpt', 'Akssgpt')

repository_path = r'C:\Users\akashdwivedi\OneDrive - Microsoft\Desktop\P_C_Ex\AKSGpt\AKSgpt'
traverse_directory(repository_path)
