-- Grant necessary permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO slothutil;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO slothutil;

-- Create views for reporting (optional)
CREATE OR REPLACE VIEW quote_stats AS
SELECT 
    category,
    length,
    COUNT(*) as total_quotes,
    COUNT(DISTINCT DATE(created_at)) as days_active,
    MAX(created_at) as last_generated
FROM quotes 
WHERE source = 'ai_generated'
GROUP BY category, length;

CREATE OR REPLACE VIEW api_usage_stats AS
SELECT 
    endpoint,
    method,
    DATE(created_at) as date,
    COUNT(*) as request_count,
    AVG(response_time_ms) as avg_response_time,
    COUNT(CASE WHEN status_code >= 400 THEN 1 END) as error_count
FROM api_usage 
GROUP BY endpoint, method, DATE(created_at)
ORDER BY date DESC;

-- Print success message
DO $$
BEGIN
    RAISE NOTICE 'Database initialized successfully for Sloth Util Cloud Agnostic Architecture';
    RAISE NOTICE 'Sample admin user: admin@slothutil.com (password: password)';
    RAISE NOTICE 'Database ready for Lambda function connections';
END $$;
`
}