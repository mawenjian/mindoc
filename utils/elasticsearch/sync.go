package elasticsearch
 
import (
  "context"
  "github.com/elastic/go-elasticsearch/v8"
  "github.com/mindoc-org/mindoc/models"
  "log"
)
 
type SyncService struct {
  esClient *elasticsearch.Client
  indexName string
}
 
// 同步单篇文档
func (s *SyncService) SyncDocument(doc *models.DocumentModel) error {
  docJSON, err := doc.ToESDocument()
  if err != nil {
    return err
  }
  
  req := esapi.IndexRequest{
    Index:      s.indexName,
    DocumentID: strconv.Itoa(doc.DocumentId),
    Body:       strings.NewReader(docJSON),
    Refresh:    "wait_for",
  }
  
  resp, err := req.Do(context.Background(), s.esClient)
  if err != nil {
    return err
  }
  defer resp.Body.Close()
  
  return nil
}
