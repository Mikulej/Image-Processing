import time
import os.path
import glob
import img_process_cv
import img_process_pil

class Timer:
    def __enter__(self):
        self.start = time.time()
        return self
    def __exit__(self, type, value, traceback):
        print(f"{(time.time() - self.start) / N * 1000} ms")

if __name__ == '__main__':
    N = 12
    dir_names = ["1k", "2k", "4k"]
    base_dir = os.path.dirname(os.path.dirname(os.getcwd()))
    out_dir = os.path.join(base_dir, "Python", "out")
    print(f"Calculating average time for {N} runs.")
    for dir_name in dir_names:
        photos_dir = os.path.join(base_dir, "source_images", dir_name)
        photos = glob.glob(os.path.join(photos_dir, "*.jpg"))
        fname1 = os.path.basename(photos[0])
        fname2 = os.path.basename(photos[1])
        print(f"Test for {dir_name} files.")
        print(f"Blur (k5) with OpenCV:")
        with Timer():
            for i in range(N):
                img_process_cv.blur(
                    photos[0], 
                    os.path.join(out_dir, f"{fname1}_blurredcv.jpg"), 
                    5
                )
        print(f"Blur (k5) with Pillow:")
        with Timer():
            for i in range(N):
                img_process_pil.blur(
                    photos[0], 
                    os.path.join(out_dir, f"{fname1}_blurredpil.jpg"), 
                    5
                )
        print(f"Blend (0.5/0.5) with OpenCV:")
        with Timer():
            for i in range(N):
                img_process_cv.blend(
                    photos[0], 
                    photos[1], 
                    os.path.join(out_dir, "blendcv.jpg"), 
                    0.5
                )
        print(f"Blend (0.5/0.5) with Pillow:")
        with Timer():
            for i in range(N):
                img_process_pil.blend(
                    photos[0], 
                    photos[1], 
                    os.path.join(out_dir, "blendpil.jpg"), 
                    0.5
                )
        print(f"Draw a circle (r50) with OpenCV:")
        with Timer():
            for i in range(N):
                img_process_cv.draw_cricle(
                    photos[0], 
                    os.path.join(out_dir, f"{fname1}_circlecv.jpg"), 
                    100, 
                    100, 
                    50
                )
        print(f"Draw a circle (r50) with Pillow:")
        with Timer():
            for i in range(N):
                img_process_pil.draw_cricle(
                    photos[0], 
                    os.path.join(out_dir, f"{fname1}_circlepil.jpg"), 
                    100, 
                    100, 
                    50
                )
    