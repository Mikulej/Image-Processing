import cv2

def blur(src_path, dst_path, ksize):
    src = cv2.imread(src_path, cv2.IMREAD_UNCHANGED)
    if src is None:
        return
    dst = cv2.GaussianBlur(src, (ksize, ksize), sigmaX=0, sigmaY=None)
    cv2.imwrite(dst_path, dst)

def blend(src1_path, src2_path, dst_path, weight):
    src1 = cv2.imread(src1_path, cv2.IMREAD_UNCHANGED)
    src2 = cv2.imread(src2_path, cv2.IMREAD_UNCHANGED)
    if src1 is None or src2 is None:
        return
    dst = cv2.addWeighted(src1, weight, src2, 1-weight, 0)
    cv2.imwrite(dst_path, dst)

if __name__ == "__main__":
    print("What would you like to do?")
    print("1 - Blur an image")
    print("2 - Blend two images")
    mode = int(input("Enter your choice: "))
    if mode == 1:
        src_path = input("Enter the name of the image: ")
        dst_path = input("Enter the name of the output image: ")
        ksize = int(input("Enter the kernel size (must be an odd integer): "))
        blur(src_path, dst_path, ksize)
    elif mode == 2:
        src1_path = input("Enter the name of first image: ")
        src2_path = input("Enter the name of second image: ")
        weight = float(input("Enter the weight of first image: "))
        dst_path = input("Enter the name of the output image: ")
        blend(src1_path, src2_path, dst_path, weight)