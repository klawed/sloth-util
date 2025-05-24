package com.slothutil.quotes.service;

import com.slothutil.quotes.repository.QuoteRepository;
import com.slothutil.quotes.service.impl.QuoteServiceImpl;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

/**
 * Unit tests for the QuoteServiceImpl class
 */
class QuoteServiceTest {

    @Mock
    private QuoteRepository quoteRepository;

    private QuoteService quoteService;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
        quoteService = new QuoteServiceImpl(quoteRepository);
    }

    @Test
    void testGetGreeting() {
        // Act
        String greeting = quoteService.getGreeting();

        // Assert
        assertEquals("Hello from Sloth Util Quote Generator!", greeting);
        
        // Verify - no repository interaction for this method
        verifyNoInteractions(quoteRepository);
    }

    @Test
    void testGenerateQuoteWithCategory() {
        // Arrange
        String category = "motivational";
        String length = "short";
        String expectedQuote = "\"The best way to predict the future is to create it.\" - Peter Drucker";
        when(quoteRepository.getRandomQuote(category, length)).thenReturn(expectedQuote);

        // Act
        String result = quoteService.generateQuote(category, length);

        // Assert
        assertEquals(expectedQuote, result);
        
        // Verify
        verify(quoteRepository, times(1)).getRandomQuote(category, length);
    }

    @Test
    void testGenerateQuoteWithoutParameters() {
        // Arrange
        String expectedQuote = "\"The best way to predict the future is to create it.\" - Peter Drucker";
        when(quoteRepository.getRandomQuote(null, null)).thenReturn(expectedQuote);

        // Act
        String result = quoteService.generateQuote(null, null);

        // Assert
        assertEquals(expectedQuote, result);
        
        // Verify
        verify(quoteRepository, times(1)).getRandomQuote(null, null);
    }
}
