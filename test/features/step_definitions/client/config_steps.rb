# frozen_string_literal: true

When('I download the {string} configuration') do |name|
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'config', "-i file:.config/#{name}.yaml", '-o file:reports/server.yaml')
  pid = spawn(cmd, %i[out err] => ['reports/config.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a configuration') do
  expect(@status.exitstatus).to eq(0)
  expect(File).to exist('reports/server.yaml')
end

Then('I should not have a configuration') do
  expect(@status.exitstatus).to eq(1)
  expect(File).not_to exist('reports/server.yaml')
end
