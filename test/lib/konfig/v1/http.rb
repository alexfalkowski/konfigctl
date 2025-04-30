# frozen_string_literal: true

module Konfig
  module V1
    class HTTP < Sinatra::Application
      before do
        content_type :json
      end

      post '/konfig.v1.Service/GetConfig' do
        halt 401, 'missing auth' if request.env['HTTP_AUTHORIZATION'].empty?

        req = JSON.parse(request.body.read)
        halt 404, 'version not found' if req['version'] == 'none'

        req.merge(data: Base64.encode64('hello')).to_json
      end

      post '/konfig.v1.Service/GetSecrets' do
        halt 401, 'missing auth' if request.env['HTTP_AUTHORIZATION'].empty?

        req = JSON.parse(request.body.read)
        secrets = req['secrets'] || []

        halt 404, 'secrets not found' if secrets.empty?

        secrets = secrets.transform_values { |v| Base64.encode64(v.to_s) }
        req['secrets'] = secrets

        req.to_json
      end
    end

    class HTTPServer < Nonnative::HTTPServer
      def initialize(service)
        app = Sinatra.new(HTTP) do
          configure do
            set :logging, false
          end
        end

        super(app, service)
      end
    end
  end
end
