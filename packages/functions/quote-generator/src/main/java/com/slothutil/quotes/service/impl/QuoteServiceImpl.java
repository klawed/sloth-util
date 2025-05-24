package com.slothutil.quotes.service.impl;

import com.slothutil.quotes.repository.QuoteRepository;
import com.slothutil.quotes.service.QuoteService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * Implementation of the QuoteService interface
 */
@Service
public class QuoteServiceImpl implements QuoteService {

    private final QuoteRepository quoteRepository;
    
    /**
     * Constructor with repository dependency
     * 
     * @param quoteRepository the repository for quotes
     */
    @Autowired
    public QuoteServiceImpl(QuoteRepository quoteRepository) {
        this.quoteRepository = quoteRepository;
    }

    /**
     * Get a simple greeting message
     * 
     * @return the greeting message
     */
    @Override
    public String getGreeting() {
        return "Hello from Sloth Util Quote Generator!";
    }

    /**
     * Generate a random quote
     * This is a placeholder for future implementation
     * 
     * @param category optional category for the quote
     * @param length optional preferred length
     * @return a random quote
     */
    @Override
    public String generateQuote(String category, String length) {
        // For now, this is just a placeholder
        // In the future, this will use AWS Bedrock and Cloudflare D1/KV
        return quoteRepository.getRandomQuote(category, length);
    }
}
