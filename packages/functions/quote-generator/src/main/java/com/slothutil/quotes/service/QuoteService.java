package com.slothutil.quotes.service;

/**
 * Service interface for generating quotes
 */
public interface QuoteService {
    
    /**
     * Get a simple greeting message
     * 
     * @return the greeting message
     */
    String getGreeting();
    
    /**
     * Generate a random quote
     * This is a placeholder for future implementation
     * 
     * @param category optional category for the quote
     * @param length optional preferred length
     * @return a random quote
     */
    String generateQuote(String category, String length);
}
