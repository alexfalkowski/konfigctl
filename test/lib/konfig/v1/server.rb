# frozen_string_literal: true

module Konfig
  module V1
    class Server < Service::Service
      def get_config(request, call)
        raise GRPC::Unauthenticated, 'missing auth' if call.metadata['authorization'].empty?
        raise GRPC::NotFound, 'version not found' if request.version == 'none'

        c = Config.new(application: request.application, version: request.version, environment: request.environment, continent: request.continent,
                       country: request.country, command: request.command, kind: request.kind, data: 'hello')
        GetConfigResponse.new(config: c)
      end

      def get_secrets(request, call)
        raise GRPC::Unauthenticated, 'missing auth' if call.metadata['authorization'].empty?

        secrets = request.secrets.to_h
        raise GRPC::NotFound, 'secrets not found' if secrets.empty?

        GetSecretsResponse.new(secrets: secrets)
      end
    end

    class GRPCServer < Nonnative::GRPCServer
      def initialize(service)
        svc = Server.new

        super(svc, service)
      end
    end
  end
end
