package com.slothutil.quotes.repository;

import com.slothutil.quotes.repository.impl.CloudflareQuoteRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.test.util.ReflectionTestUtils;

import static org.junit.jupiter.api.Assertions.*;

/**
 * Unit tests for the CloudflareQuoteRepository class
 */
class QuoteRepositoryTest {

    private QuoteRepository quoteRepository;

    @BeforeEach
    void setUp() {
        quoteRepository = new CloudflareQuoteRepository();
        
        // Set test values for the properties
        ReflectionTestUtils.setField(quoteRepository, "cloudflareAccountId", "test-account-id");
        ReflectionTestUtils.setField(quoteRepository, "cloudflareApiToken", "test-api-token");
        ReflectionTestUtils.setField(quoteRepository, "cloudflareKvNamespaceId", "test-kv-namespace");
        ReflectionTestUtils.setField(quoteRepository, "cloudflareD1DatabaseId", "test-d1-database");
        ReflectionTestUtils.setField(quoteRepository, "cloudflareD1WorkerUrl", "http://localhost:8787");
    }

    @Test
    void testGetRandomQuote() {
        // Act
        String quote = quoteRepository.getRandomQuote(null, null);
        
        // Assert
        assertNotNull(quote);
        assertFalse(quote.isEmpty());
        assertTrue(quote.contains("The best way to predict the future"));
        assertTrue(quote.contains("Peter Drucker"));
    }

    @Test
    void testGetRandomQuoteWithParameters() {
        // Act
        String quote = quoteRepository.getRandomQuote("motivational", "short");
        
        // Assert
        assertNotNull(quote);
        assertFalse(quote.isEmpty());
        assertTrue(quote.contains("The best way to predict the future"));
        assertTrue(quote.contains("Peter Drucker"));
    }

    @Test
    void testSaveQuote() {
        // Act
        boolean result = quoteRepository.saveQuote(
            "Life is what happens when you're busy making other plans.",
            "John Lennon",
            "life"
        );
        
        // Assert
        assertTrue(result);
    }
}
