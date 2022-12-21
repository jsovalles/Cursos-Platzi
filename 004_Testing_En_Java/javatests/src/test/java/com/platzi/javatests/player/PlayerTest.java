package com.platzi.javatests.player;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnitRunner;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;
import static org.mockito.Mockito.when;

@RunWith(MockitoJUnitRunner.class)
public class PlayerTest {

    @Mock
    Dice dice;

    @InjectMocks
    Player player;

    @Test
    public void playerLoses_whenNumberIsTooLow() {
        when(dice.roll()).thenReturn(2);
        player = new Player(dice, 3);
        assertFalse(player.play());
    }

    @Test
    public void playerWins_whenNumberIsEqualOrHigh() {
        when(dice.roll()).thenReturn(5);
        player = new Player(dice, 4);
        assertTrue(player.play());
    }

}