package com.teamProject.shapes;

import com.teamProject.Shape;

public class RectangleShape implements Shape {
    private final int x;
    private final int y;
    private final int width;
    private final int height;
    private int currentX;
    private int currentY;

    public RectangleShape(int x, int y, int width, int height) {
        this.height = height;
        this.width = width;
        this.x = x;
        this.y = y;
    }

    @Override
    public int[] reset() {
        currentX = x;
        currentY = y;
        return new int[]{x, y};
    }

    @Override
    public int[] next() {
        ++currentX;
        if (currentX >= x + width) {
            currentX = x;
            ++currentY;
        }
        if (currentY >= y + height) {
            return null;
        }
        return new int[]{currentX, currentY};
    }
}
