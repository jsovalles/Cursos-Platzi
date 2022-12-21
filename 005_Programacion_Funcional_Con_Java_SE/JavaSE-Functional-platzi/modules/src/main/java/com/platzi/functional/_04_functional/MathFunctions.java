package com.platzi.functional._04_functional;

import java.util.function.Function;
import java.util.function.Predicate;

public class MathFunctions {
    public static void main(String[] args) {
        Function<Integer, Integer> square = x -> x * x;
        System.out.println(square.apply(5));

        Predicate<Integer> isEven = x -> x % 2 == 0;
        System.out.println(isEven.test(6));

    }
}
