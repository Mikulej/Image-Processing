from PIL import Image, ImageFilter

def blur(src_path, dst_path, ksize):
    src = Image.open(src_path)
    if src is None:
        return
    dst = src.filter(ImageFilter.GaussianBlur(ksize))
    dst.save(dst_path)

def blend(src1_path, src2_path, dst_path, weight):
    src1 = Image.open(src1_path)
    src2 = Image.open(src2_path)
    if src1 is None or src2 is None:
        return
    dst = Image.blend(src1, src2, weight)
    dst.save(dst_path)

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