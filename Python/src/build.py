import PyInstaller.__main__

if __name__ == '__main__':
    PyInstaller.__main__.run([
        'main.py',
        '--noconfirm',
        '--onefile',
        '--console',
        '--distpath',
        '',
        '--name',
        'img_process'
    ])