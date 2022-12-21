package com.platzi.functional._06_reference_operator;

import java.util.Arrays;
import java.util.List;
import java.util.function.Consumer;

public class NombresUtils {
    static <T> List<T> getList(T... elements) {
        return Arrays.asList(elements);
    }

    public static void main(String[] args) {
        List<String> professors = getList("Nicolas", "Jose", "Zulma");

        Consumer<String> printer = text -> System.out.println(text);
        professors.forEach(printer);

        professors.forEach(System.out::println);

    }
}
