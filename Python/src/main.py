from img_process_pil import blur, blend, draw_cricle, draw_rectangle
#from img_process_cv import blur, blend, draw_cricle, draw_rectangle

if __name__ == "__main__":
    print("1 - Blur an image")
    print("2 - Blend two images")
    print("3 - Draw a circle")
    print("4 - Draw a rectangle")
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
    elif mode == 3:
        src_path = input("Enter the name of the image: ")
        dst_path = input("Enter the name of the output image: ")
        x = int(input("Enter the x coordinate of the center: "))
        y = int(input("Enter the y coordinate of the center: "))
        r = int(input("Enter the radius: "))
        draw_cricle(src_path, dst_path, x, y, r)
    elif mode == 4:
        src_path = input("Enter the name of the image: ")
        dst_path = input("Enter the name of the output image: ")
        x1 = int(input("Enter the x coordinate of the first corner: "))
        y1 = int(input("Enter the y coordinate of the first corner: "))
        x2 = int(input("Enter the x coordinate of the second corner: "))
        y2 = int(input("Enter the y coordinate of the second corner: "))
        draw_rectangle(src_path, dst_path, x1, y1, x2, y2)