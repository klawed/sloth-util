package com.slothutil.quotes;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.LambdaLogger;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;
import org.junit.jupiter.api.Tag;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.test.context.TestPropertySource;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.doNothing;
import static org.mockito.Mockito.when;

/**
 * Integration tests for the QuoteGenerator
 * These tests verify the entire flow from handler to repository
 */
@ExtendWith(MockitoExtension.class)
@Tag("integration")
@TestPropertySource(properties = {
    "ARCHITECTURE_TYPE=hybrid",
    "CLOUDFLARE_ACCOUNT_ID=test-account-id",
    "CLOUDFLARE_API_TOKEN=test-api-token",
    "CLOUDFLARE_KV_NAMESPACE_ID_QUOTES=test-kv-namespace",
    "CLOUDFLARE_D1_DATABASE_ID=test-d1-database",
    "CLOUDFLARE_D1_WORKER_URL=http://localhost:8787"
})
class QuoteGeneratorIntegrationTest {

    @Mock
    private Context context;

    @Mock
    private LambdaLogger logger;

    /**
     * Test the complete flow of the quote generator with a GET request
     */
    @Test
    void testQuoteGeneratorFlow() {
        // Arrange
        when(context.getLogger()).thenReturn(logger);
        doNothing().when(logger).log(String.valueOf(org.mockito.ArgumentMatchers.any()));
        
        // Create the handler directly (Spring context will be initialized)
        QuoteHandler handler = new QuoteHandler();
        
        // Create a request
        APIGatewayProxyRequestEvent request = new APIGatewayProxyRequestEvent();
        request.setHttpMethod("GET");
        request.setPath("/quotes/random");
        
        // Act
        APIGatewayProxyResponseEvent response = handler.handleRequest(request, context);
        
        // Assert
        assertEquals(200, response.getStatusCode());
        assertNotNull(response.getBody());
        assertTrue(response.getBody().contains("message"));
        assertTrue(response.getBody().contains("Hello from Sloth Util Quote Generator!"));
        assertTrue(response.getBody().contains("timestamp"));
        
        // Verify headers
        assertEquals("application/json", response.getHeaders().get("Content-Type"));
        assertEquals("*", response.getHeaders().get("Access-Control-Allow-Origin"));
    }
}
