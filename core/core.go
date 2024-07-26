package core

import (
   "context"
   "fmt"
   "net/url"
   "strings"

   "github.com/go-logr/logr"
)

type URLProvider interface {
   GetURL() (string, error)
}

func ExtractURL(ctx context.Context, provider URLProvider) error {
   logger := logr.FromContextOrDiscard(ctx)
   logger.V(1).Info("Debug: Entering ExtractURL function")

   content, err := provider.GetURL()
   if err != nil {
   	return fmt.Errorf("failed to get URL: %w", err)
   }

   return ParseAndPrintURL(ctx, content)
}

func ParseAndPrintURL(ctx context.Context, content string) error {
   logger := logr.FromContextOrDiscard(ctx)

   content = strings.TrimSpace(content)

   parsedURL, err := url.Parse(content)
   if err != nil {
   	return fmt.Errorf("invalid URL: %w", err)
   }

   fmt.Printf("Scheme: %s\n", parsedURL.Scheme)
   fmt.Printf("Host: %s\n", parsedURL.Host)
   fmt.Printf("Path: %s\n", parsedURL.Path)

   pathComponents := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
   for i, component := range pathComponents {
   	fmt.Printf("Path Component %d: %s\n", i+1, component)
   }

   if strings.Contains(parsedURL.Path, "/maps/") {
   	dataIndex := -1
   	for i, component := range pathComponents {
   		if component == "data" {
   			dataIndex = i
   			break
   		}
   	}

   	if dataIndex != -1 && dataIndex+1 < len(pathComponents) {
   		dataParams := strings.Split(pathComponents[dataIndex+1], "!")
   		for _, param := range dataParams {
   			parts := strings.SplitN(param, "=", 2)
   			if len(parts) == 2 {
   				fmt.Printf("Data Param: %s = %s\n", parts[0], parts[1])
   			} else {
   				fmt.Printf("Data Param: %s\n", parts[0])
   			}
   		}
   	}
   }

   queryParams := parsedURL.Query()
   for key, values := range queryParams {
   	for _, value := range values {
   		fmt.Printf("Query Param: %s = %s\n", key, value)
   	}
   }

   if parsedURL.Fragment != "" {
   	fmt.Printf("Fragment: %s\n", parsedURL.Fragment)
   }

   logger.V(1).Info("Debug: Exiting ParseAndPrintURL function")

   return nil
}
