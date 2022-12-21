package com.platzi.javatests.utils;

import org.junit.Assert;
import org.junit.Test;


public class StringUtilsTest {

    @Test
    public void testRepeat() {
        String result = StringUtils.repeat("hola", 3);
        Assert.assertEquals("holaholahola", result);
    }

    @Test
    public void testRepeatZeroTimes() {
        String result = StringUtils.repeat("hola", 0);
        Assert.assertEquals("", result);
    }

    @Test(expected = IllegalArgumentException.class)
    public void testRepeatNegativeTimes() {
        String result = StringUtils.repeat("hola", -1);
        Assert.assertEquals("", result);
    }

    @Test
    public void testIfStringIsNotNull(){
        Boolean result = StringUtils.isEmpty("hola");
        Assert.assertFalse(result);
    }

    @Test
    public void testIfStringIsEmpty(){
        Boolean result = StringUtils.isEmpty("");
        Assert.assertTrue(result);
    }

    @Test
    public void testIfNullIsEmpty(){
        Boolean result = StringUtils.isEmpty(null);
        Assert.assertTrue(result);
    }

    @Test
    public void testIfSpaceIsEmpty(){
        Boolean result = StringUtils.isEmpty(" ");
        Assert.assertTrue(result);
    }
}