package com.teamProject.processors;

import com.teamProject.ImageProcessor;
import com.teamProject.Shape;

public class RemoveProcessor extends ImageProcessor {
    static final int R = 0;
    static final int G = 1;
    static final int B = 2;

    Shape[] shapes;
    int[] defaultColor;

    public RemoveProcessor(Shape[] shapes, int[] defaultColor) {
        this.shapes = shapes;
        this.defaultColor = defaultColor;
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

        for (Shape shape : shapes) {
            int[] coords = shape.reset();
            while (coords != null) {
                pixels[toIndex(coords[0], coords[1])] = toPixelValue(new int[]{defaultColor[R], defaultColor[G], defaultColor[B]});
                coords = shape.next();
            }
        }

        return pixels;
    }

    @Override
    public String getName() {
        return "Remove (" + shapes.length + " shapes)";
    }
}
