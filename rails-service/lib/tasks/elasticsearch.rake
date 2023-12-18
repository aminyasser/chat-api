namespace :elasticsearch do
    desc "Create Elasticsearch index for Messages"
    task :create_index => :environment do
        # Message.reindex
        Message.__elasticsearch__.create_index!
        Message.__elasticsearch__.refresh_index!
        Message.import
    end
end