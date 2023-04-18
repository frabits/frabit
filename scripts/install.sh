# Frabit - The next-generation database automatic operation platform
# Copyright © 2022-2023 Blylei <blylei.info@gmail.com>
#
# Licensed under the GNU General Public License, Version 3.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#	https://www.gnu.org/licenses/gpl-3.0.txt
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

abort() {
    printf "%s\n" "$@" >&2
    exit 1
}

apiURL="https://api.github.com/repos/frabits/frabit"
packageURL="https://github.com/frabits/frabit/releases/download"
tmp_dir="/tmp/frabit"
base_dir="/usr/local/frabit"
install_dir="/usr/local/bin"

prepare_dir(){
  mkdir -p $tmp_dir || abort "$tmp_dir create failed"
  mkdir -p $base_dir || abort "$base_dir create failed"
  mkdir -p $install_dir || abort "$install_dir create failed"
  return 0
}

get_latest(){
  local latestURL=$apiURL"/releases/latest"
  local version=$(curl -sX GET "$latestURL" | awk '/tag_name/{print $4;exit}' FS='[""]')
  echo "$version"
  reture 0
}

detect_arch(){
  local allowed_arch="x86_64 arm64"
  local arch=$(uname -m)

  for i in $allowed_arch;
  do
    if [ "$i" == "$arch" ];then
      echo "$arch"
      reture 0
    fi
  done
  abort "Arch ${arch} is not support, frabit is only supported on x86_64 and arm64"
}

detect_dist(){
  local allowed_dist="linux darwin"
  local dist=$(uname -s)

  for i in $allowed_dist;
  do
    if [ "$i" == "$dist" ];then
      echo "$dist"
      reture 0
    fi
  done
  abort "Arch ${dist} is not support, frabit is only supported on Linux and darwin/MacOS"
}

detect_os(){
  arch=$(detect_arch)
  dist=$(detect_dist)
  echo "$dist"_"$arch"
}

test_curl() {
    local curl_version=$(curl --version 2>/dev/null)
    if [ $? -ne 0 ]; then
        abort "You must install curl before installing frabit."
    fi
}

test_tar() {
    local tar_version=$(tar --version 2>/dev/null)
    if [ $? -ne 0 ]; then
        abort "You must install tar before installing frabit."
    fi
}

# package uri: https://github.com/frabits/frabit/releases/download/v2.1.17/frabit_2.1.17_linux_arm64.tar.gz
download_package() {
    local local_file=$1
    local source_url=$2
    echo "Downloading ${source_url}..."

    local code=$(curl -w '%{http_code}' -L -o "${local_file}" "${source_url}")
    if [ "$code" != "200" ]; then
        abort "Failed to download from ${source_url}, status code: ${code}"
    fi

    echo "Downloading ${source_url} successes"
}

gen_config(){
  # /etc/frabit/frabit.yml
  echo ""
}

gen_server_service_file(){
  # /usr/lib/systemd//frabit-server.service
  echo ""
}

gen_agent_service_file(){
  # /usr/lib/systemd//frabit-agent.service
  echo ""
}

purge(){
  echo ""
}

pre_check(){
  test_curl
  test_tar
  prepare_dir
}

post_check(){
  gen_config
  gen_server_service_file
  gen_agent_service_file

  purge
}

main(){
    pre_check
    local version=$(get_latest)
    local os=$(detect_os)
    local package=frabit_"$version"_"$os".tar.gz
    download_package $tmp_dir/"$package" $packageURL/"$version"/"$package"
    post_check
}

main