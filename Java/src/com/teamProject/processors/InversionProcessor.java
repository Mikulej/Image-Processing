package com.teamProject.processors;

import com.teamProject.ImageProcessor;

public class InversionProcessor extends ImageProcessor {
    static final int R = 0;
    static final int G = 1;
    static final int B = 2;

    @Override
    public int[] process(int[] pixels, int width, int height) {
        for (int i = 0; i < pixels.length; ++i) {
            int[] channels = {(pixels[i] >> 16) & 255, (pixels[i] >> 8) & 255, pixels[i] & 255};
            channels[R] = 255 - channels[R];
            channels[G] = 255 - channels[G];
            channels[B] = 255 - channels[B];
            pixels[i] = 0xFF000000 + channels[R] * 256 * 256 + channels[G] * 256 + channels[B];
        }

        return pixels;
    }

    @Override
    public String getName() {
        return "Inversion";
    }
}
