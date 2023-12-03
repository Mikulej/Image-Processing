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
        print(f"Took {time.time() - self.start} s")


if __name__ == '__main__':
    N = 12
    photos_dir = os.path.dirname(os.getcwd())
    photos_dir = os.path.join(photos_dir, "source_images", "4k")
    photos = glob.glob(os.path.join(photos_dir, "*.jpg"))
    fname1 = os.path.basename(photos[0])
    fname2 = os.path.basename(photos[1])
    print(f"Blur {fname1} {N} times with OpenCV")
    with Timer():
        for i in range(N):
            img_process_cv.blur(photos[0], os.path.join("out", f"{fname1}_blurredcv.jpg"), 5)
    print(f"Blur {fname1} {N} times with Pillow")
    with Timer():
        for i in range(N):
            img_process_pil.blur(photos[0], os.path.join("out", f"{fname1}_blurredpil.jpg"), 5)
    print(f"Blend {fname1} and {fname2} {N} times with OpenCV")
    with Timer():
        for i in range(N):
            img_process_cv.blend(photos[0], photos[1], os.path.join("out", "blendcv.jpg"), 0.5)
    print(f"Blend {fname1} and {fname2} {N} times with Pillow")
    with Timer():
        for i in range(N):
            img_process_pil.blend(photos[0], photos[1], os.path.join("out", "blendpil.jpg"), 0.5)