package com.torabian.fireback.firebackspring;

public class BinarySearch {
    int binarySearch(int items[], int x) {
        int left = 0, right = items.length - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (items[mid] == x) {
                return mid;
            }

            if (items[mid] < x) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }

        }

        return -1;
    }
}
