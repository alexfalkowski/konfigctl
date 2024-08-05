# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'konfig/v1/service_services_pb'
require 'konfig/v1/server'

module Konfig
  class << self
    def client_config
      @client_config ||= Nonnative.configurations('.config/client.yml')
    end
  end
end
