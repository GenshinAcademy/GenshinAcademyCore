#!/bin/bash

#######################################
# Defining color codes
#######################################
readonly GREEN='\e[32m'
readonly BLUE='\e[34m'
readonly YELLOW='\e[33m'
readonly RED='\e[31m'
readonly NC='\e[0m' # No Color

#######################################
# Configuring variables for script
#######################################

readonly LANGUAGES=("English" "Russian")

readonly BASE_URL="https://api.github.com/repos/theBowja/genshin-db/contents/src/data"

#######################################
# Functions
#######################################

#######################################
# Prints error message
# Arguments:
#   String
#######################################
function print_error () {
    printf "${RED}Error: %s${NC}\n" "${1}"
}

#######################################
# Prints info message
# Arguments:
#   String
#######################################
function print_info () {
    printf "${GREEN}%s${NC}\n" "${1}"
}

#######################################
# Prints status message
# Arguments:
#   String
#######################################
function print_status () {
    printf "${BLUE}%s${NC}\n" "${1}"
}

#######################################
# Prints input message
# Arguments:
#   String
#######################################
function print_input () {
    printf "${YELLOW}%s${NC}\n" "${1}"
}

#######################################
# Reads input string
# Arguments:
#   Message to display
#   Variable to store an input
#######################################
function get_input_string () {
    printf "${YELLOW}"
    read -rp "${1}" "${2}"
    printf "${NC}\n"
}
#######################################
# Reads input as Y/N
# Arguments:
#   Message to display
#   Variable to store an input
#######################################
function get_input_bool () {
    print_input "${1} (y/n)"
    read -rsn1 "${2}"
}

#######################################
# Prints curl response output
# Arguments:
#   Response from curl
#######################################
function print_response () {
    printf "${YELLOW}\nRESPONSE:\n%s${NC}\n\n\n" "${1}"
}

#######################################
# Checks whether curl was ok
# Arguments:
#   Variable from curl result to check
#######################################
function check_curl () {
    if [[ -z ${1} || ${1} = "null" ]]
    then
        print_error "curl request failed"
        exit 1
    fi
}

#######################################
# Installs required packages
# Arguments:
#   None
#######################################
function install_dependencies () {
    print_status "Looking for required dependencies."
    
    packages=(curl jq wget)
    for pkg in "${packages[@]}"; do
        if ! dpkg-query -W -f='${Status}' "${pkg}" | grep -q "install ok installed"; then
            sudo apt-get update
            break
        fi
    done
    
    for pkg in "${packages[@]}"; do
        if ! dpkg-query -W -f='${Status}' "${pkg}" | grep -q "install ok installed"; then
            print_error "${pkg} package not found."
            print_status "Installing ${pkg} package."
            sudo apt-get install -y "${pkg}" || {
                print_error "Could not install ${pkg} package."
                exit 1
            }
        else
            print_info "${pkg} package already installed."
        fi
    done
}

function init() {
    install_dependencies
    
    get_input_string "Enter Github Token: " TOKEN
}

function main() {
    for lang in "${LANGUAGES[@]}"; do
        
        print_status "Parsing ${lang} data..."

        url="${BASE_URL}/${lang}/characters/"
        
        response=$(curl --location "${url}" -H "Authorization: Bearer ${TOKEN}")

        print_response "${response}"

        # Parse files list from github repository folder
        files=$(printf "%s" "${response}" | jq -r '.[].download_url')
        
        check_curl "${files}"

        # Download files from parsed list
        for file in ${files}; do
            print_status "Downloading ${file}..."
            wget --header="Authorization: Bearer ${TOKEN}" -P "data/${lang}" -b -q "${file}"
        done
    done
}

init
main