import PyInstaller.__main__

if __name__ == '__main__':
    PyInstaller.__main__.run([
        'img_process_pil.py',
        '--noconfirm',
         '--onefile',
        '--console',
        '--distpath',
        ''
    ])
    PyInstaller.__main__.run([
        'img_process_cv.py',
        '--noconfirm',
         '--onefile',
        '--console',
        '--distpath',
        ''
    ])