package com.slothutil.quotes.repository;

/**
 * Repository interface for accessing quotes
 */
public interface QuoteRepository {
    
    /**
     * Get a random quote from the repository
     * 
     * @param category optional category filter
     * @param length optional length preference
     * @return a random quote string
     */
    String getRandomQuote(String category, String length);
    
    /**
     * Save a quote to the repository
     * 
     * @param quote the quote text
     * @param author the quote author
     * @param category the quote category
     * @return true if saved successfully, false otherwise
     */
    boolean saveQuote(String quote, String author, String category);
}
