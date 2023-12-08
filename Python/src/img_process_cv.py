import cv2

def blur(src_path: str, dst_path: str, ksize: int) -> None:
    try:
        src = cv2.imread(src_path, cv2.IMREAD_UNCHANGED)
    except:
        print(f"Error reading {src_path}")
        return
    dst = cv2.GaussianBlur(src, (ksize, ksize), sigmaX=0, sigmaY=None)
    cv2.imwrite(dst_path, dst)

def blend(src1_path: str, src2_path: str, dst_path: str, weight: float) -> None:
    try:
        src1 = cv2.imread(src1_path, cv2.IMREAD_UNCHANGED)
    except:
        print(f"Error reading {src1_path}")
        return
    try:
        src2 = cv2.imread(src2_path, cv2.IMREAD_UNCHANGED)
    except:
        print(f"Error reading {src2_path}")
        return
    dst = cv2.addWeighted(src1, weight, src2, 1-weight, 0)
    cv2.imwrite(dst_path, dst)

def draw_cricle(src_path: str, dst_path: str, x: int, y: int, r: int) -> None:
    try:
        src = cv2.imread(src_path, cv2.IMREAD_UNCHANGED)
    except:
        print(f"Error reading {src_path}")
        return
    dst = src.copy()
    cv2.circle(dst, (x, y), r, (255, 255, 255), -1)
    cv2.imwrite(dst_path, dst)

def draw_rectangle(src_path: str, dst_path: str, x1: int, y1: int, x2: int, y2: int) -> None:
    try:
        src = cv2.imread(src_path, cv2.IMREAD_UNCHANGED)
    except:
        print(f"Error reading {src_path}")
        return
    dst = src.copy()
    cv2.rectangle(dst, (x1, y1), (x2, y2), (255, 255, 255), -1)
    cv2.imwrite(dst_path, dst)    
