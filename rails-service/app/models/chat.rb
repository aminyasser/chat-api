class Chat < ApplicationRecord
    has_many :message
    belongs_to :app , foreign_key: 'app_token', primary_key: 'app_token', optional: true

    validates :app_token, :presence => true
end
