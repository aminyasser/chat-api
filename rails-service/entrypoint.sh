#  bundle exec rake install
bundle exec rake elasticsearch:create_index
bundle exec rails server -b 0.0.0.0