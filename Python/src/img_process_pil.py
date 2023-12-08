from PIL import Image, ImageFilter, ImageDraw

def blur(src_path: str, dst_path: str, ksize: int) -> None:
    try:
        src = Image.open(src_path)
    except:
        print(f"Error reading {src_path}")
        return
    dst = src.filter(ImageFilter.GaussianBlur(ksize))
    dst.save(dst_path)

def blend(src1_path: str, src2_path: str, dst_path: str, weight: float) -> None:
    try:
        src1 = Image.open(src1_path)
    except:
        print(f"Error reading {src1_path}")
        return
    try:
        src2 = Image.open(src2_path)
    except:
        print(f"Error reading {src2_path}")
        return
    dst = Image.blend(src1, src2, weight)
    dst.save(dst_path)

def draw_cricle(src_path: str, dst_path: str, x: int, y: int, r: int) -> None:
    try:
        src = Image.open(src_path)
    except:
        print(f"Error reading {src_path}")
        return
    dst = src.copy()
    draw = ImageDraw.Draw(dst)
    draw.ellipse((x-r, y-r, x+r, y+r), fill=(255,255,255))
    dst.save(dst_path)

def draw_rectangle(src_path: str, dst_path: str, x1: int, y1: int, x2: int, y2: int) -> None:
    try:
        src = Image.open(src_path)
    except:
        print(f"Error reading {src_path}")
        return
    dst = src.copy()
    draw = ImageDraw.Draw(dst)
    draw.rectangle((x1, y1, x2, y2), fill=(255,255,255))
    dst.save(dst_path)
