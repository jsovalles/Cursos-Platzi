package com.platzi.javatests.utils;

public class StringUtils {

    public static String repeat(String str, int times) {

        if (times < 0)
            throw new IllegalArgumentException("hehe");

        String result = "";

        for (int i = 0; i < times; i++) {
            result += str;
        }
        return result;
    }


    public static boolean isEmpty(String str) {
        if(str == null || str.trim().length() <= 0)
            return true;
        return false;
    }

}
