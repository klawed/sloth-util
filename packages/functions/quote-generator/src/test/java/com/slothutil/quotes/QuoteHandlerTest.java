package com.slothutil.quotes;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.LambdaLogger;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;
import com.slothutil.quotes.service.QuoteService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

/**
 * Unit tests for the QuoteHandler class
 */
class QuoteHandlerTest {

    @Mock
    private QuoteService quoteService;

    @Mock
    private Context context;

    @Mock
    private LambdaLogger logger;

    private QuoteHandler quoteHandler;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
        when(context.getLogger()).thenReturn(logger);
        doNothing().when(logger).log(anyString());
        quoteHandler = new QuoteHandler(quoteService);
    }

    @Test
    void testHandleRequestSuccess() {
        // Arrange
        APIGatewayProxyRequestEvent request = new APIGatewayProxyRequestEvent();
        String expectedGreeting = "Hello from Sloth Util Quote Generator!";
        when(quoteService.getGreeting()).thenReturn(expectedGreeting);

        // Act
        APIGatewayProxyResponseEvent response = quoteHandler.handleRequest(request, context);

        // Assert
        assertEquals(200, response.getStatusCode());
        assertTrue(response.getBody().contains(expectedGreeting));
        assertEquals("application/json", response.getHeaders().get("Content-Type"));
        assertEquals("*", response.getHeaders().get("Access-Control-Allow-Origin"));
        
        // Verify
        verify(quoteService, times(1)).getGreeting();
        verify(context, times(1)).getLogger();
        verify(logger, times(1)).log(anyString());
    }

    @Test
    void testHandleRequestException() {
        // Arrange
        APIGatewayProxyRequestEvent request = new APIGatewayProxyRequestEvent();
        when(quoteService.getGreeting()).thenThrow(new RuntimeException("Test exception"));

        // Act
        APIGatewayProxyResponseEvent response = quoteHandler.handleRequest(request, context);

        // Assert
        assertEquals(500, response.getStatusCode());
        assertEquals("{\"error\":\"Internal Server Error\"}", response.getBody());
        
        // Verify
        verify(quoteService, times(1)).getGreeting();
        verify(context, times(2)).getLogger();  // Once for initial log, once for error
        verify(logger, times(2)).log(anyString());
    }
}
