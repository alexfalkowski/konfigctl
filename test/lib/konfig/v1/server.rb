# frozen_string_literal: true

module Konfig
  module V1
    class Server < Service::Service
      def get_config(request, _call)
        c = Config.new(application: request.application, version: request.version, environment: request.environment, continent: request.continent,
                       country: request.country, command: request.command, kind: request.kind, data: 'hello')
        GetConfigResponse.new(config: c)
      end

      def get_secrets(request, _call)
        GetSecretsResponse.new(secrets: request.secrets.to_h)
      end
    end

    class GRPCServer < Nonnative::GRPCServer
      def svc
        Server.new
      end
    end
  end
end
