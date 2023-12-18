# Start from the Ruby 3.0.2 Alpine image
FROM ruby:3.0.2

# Install dependencies
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs


# Set the working directory
WORKDIR /app

# Copy the Gemfile and Gemfile.lock
COPY Gemfile Gemfile.lock ./

# Install bundler and gems
RUN gem install bundler -i '2.2.22'

RUN bundle install

# Copy the current directory contents into the container
COPY . ./

# Expose port 3000
EXPOSE 3000

# Start the main process.
CMD ["rails", "server", "-b", "0.0.0.0"]