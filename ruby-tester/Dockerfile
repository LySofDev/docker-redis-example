FROM ruby:2.5.3-alpine3.8
RUN gem install bundle
RUN mkdir /ruby
COPY Gemfile* /ruby/
WORKDIR /ruby
RUN bundle
COPY . /ruby/
