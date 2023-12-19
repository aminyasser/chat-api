#  bundle exec rake install
rm -f tmp/pids/server.pid
bundle exec rake elasticsearch:create_index
bundle exec rails db:migrate
bundle exec rails server -b 0.0.0.0