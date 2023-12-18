class App < ApplicationRecord
    has_secure_token :app_token, length: 36

    validates :name, :presence => true
end
