
class Message < ApplicationRecord
    include Elasticsearch::Model
    include Elasticsearch::Model::Callbacks

    belongs_to :chat
    validates :body, :presence => true


    index_name "messages"

    settings index: { number_of_shards: 1 } do
        mappings dynamic: 'false' do
            indexes :chat_id, type: 'keyword'
            indexes :body, type: 'text', analyzer: 'standard'
        end
        end

        def as_indexed_json(options={})
        as_json(only: [:chat_id, :body])
    end
        
    def self.search(query, chat_id)
        search_definition = {
          query: {
            bool: {
              must: {
                multi_match: {
                  query: query,
                  fields: ['body']
                }
              }
            }
          }
        }
        
        # Filter by chat_id if provided
        search_definition[:query][:bool][:filter] = { term: { chat_id: chat_id } } if chat_id
      
        __elasticsearch__.search(search_definition).records.to_json(except: [:id, :chat_id, created_at: :updated_at])
      end
end
