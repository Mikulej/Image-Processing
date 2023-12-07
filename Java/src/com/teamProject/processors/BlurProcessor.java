package com.teamProject.processors;

import com.teamProject.ImageProcessor;

public class BlurProcessor extends ImageProcessor {
    static final int R = 0;
    static final int G = 1;
    static final int B = 2;

    int boxSize;

    public BlurProcessor(int boxSize) {
        this.boxSize = boxSize;
        if (boxSize % 2 == 0) {
            ++this.boxSize;
        }
    }

    private int toIndex(int x, int y) {
        return x + width * y;
    }

    private int[] fromPixelValue(int pixelValue) {
        return new int[]{(pixelValue >> 16) & 255, (pixelValue >> 8) & 255, pixelValue & 255};
    }

    private int toPixelValue(int[] channels) {
        return 0xFF000000 + channels[R] * 256 * 256 + channels[G] * 256 + channels[B];
    }

    @Override
    public int[] process(int[] pixels, int width, int height) {
        super.process(pixels, width, height);

        int[] newPixels = new int[pixels.length];

        for (int y = 0; y < height; ++y) {
            for (int x = 0; x < width; ++x) {
                int[] avg = new int[3];
                for (int i = 0; i < boxSize * boxSize; ++i) {
                    int newX = x + (i % boxSize) - boxSize / 2;
                    int newY = y + (i / boxSize) - boxSize / 2;

                    if (newX < 0 || newX >= width || newY < 0 || newY >= height) {
                        continue;
                    }

                    avg[R] += fromPixelValue(pixels[toIndex(newX, newY)])[R];
                    avg[G] += fromPixelValue(pixels[toIndex(newX, newY)])[G];
                    avg[B] += fromPixelValue(pixels[toIndex(newX, newY)])[B];
                }
                avg[R] = avg[R] / (boxSize * boxSize);
                avg[G] = avg[G] / (boxSize * boxSize);
                avg[B] = avg[B] / (boxSize * boxSize);
                newPixels[toIndex(x, y)] = toPixelValue(avg);
            }
        }

        return newPixels;
    }

    @Override
    public String getName() {
        return "Box blur (" + boxSize + "x" + boxSize + ")";
    }
}
