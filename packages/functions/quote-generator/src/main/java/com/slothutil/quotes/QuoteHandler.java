package com.slothutil.quotes;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;
import com.slothutil.quotes.service.QuoteService;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

import java.util.HashMap;
import java.util.Map;

/**
 * AWS Lambda function handler for the Quote Generator
 * Responds with a simple greeting for now
 */
@Configuration
@ComponentScan(basePackages = "com.slothutil.quotes")
public class QuoteHandler implements RequestHandler<APIGatewayProxyRequestEvent, APIGatewayProxyResponseEvent> {

    private final QuoteService quoteService;

    /**
     * Constructor used by AWS Lambda runtime
     */
    public QuoteHandler() {
        // Initialize Spring context
        AnnotationConfigApplicationContext context = new AnnotationConfigApplicationContext();
        context.register(QuoteHandler.class);
        context.refresh();
        
        // Get the QuoteService bean
        this.quoteService = context.getBean(QuoteService.class);
    }

    /**
     * Constructor for testing with dependency injection
     *
     * @param quoteService the quote service implementation
     */
    public QuoteHandler(QuoteService quoteService) {
        this.quoteService = quoteService;
    }

    /**
     * Handle the API Gateway request for quote generation
     *
     * @param request the API Gateway request event
     * @param context the AWS Lambda context
     * @return the API Gateway response event
     */
    @Override
    public APIGatewayProxyResponseEvent handleRequest(APIGatewayProxyRequestEvent request, Context context) {
        try {
            context.getLogger().log("Processing quote request");
            
            // Get the greeting from the service
            String greeting = quoteService.getGreeting();

            // Prepare response with proper headers
            Map<String, String> headers = new HashMap<>();
            headers.put("Content-Type", "application/json");
            headers.put("Access-Control-Allow-Origin", "*");
            headers.put("Access-Control-Allow-Methods", "GET, OPTIONS");
            headers.put("Access-Control-Allow-Headers", "Content-Type,Authorization");

            // Return the response
            return new APIGatewayProxyResponseEvent()
                    .withStatusCode(200)
                    .withHeaders(headers)
                    .withBody(String.format("{\"message\":\"%s\",\"timestamp\":\"%s\"}", 
                            greeting, 
                            java.time.Instant.now().toString()));
        } catch (Exception e) {
            context.getLogger().log("Error: " + e.getMessage());
            
            // Return error response
            return new APIGatewayProxyResponseEvent()
                    .withStatusCode(500)
                    .withBody("{\"error\":\"Internal Server Error\"}");
        }
    }
}
