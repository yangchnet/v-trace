env_file="$1"
template_file="$2"
target_file="$3"

if [ $# -ne 3 ];then
    >&2 echo  "Usage: genconfig.sh <env_file> <template> <target>"
    exit 1
fi

source "${env_file}"

declare -A envs

set +u
for env in $(sed -n 's/^[^#].*${\(.*\)}.*/\1/p' ${template_file})
do
    if [ -z "$(eval echo \$${env})" ];then
        >&2 echo "environment variable '${env}' not set"
    fi
done

if [ "${missing}" ];then
    exit 1
fi

eval "cat << EOF
$(cat ${template_file})
EOF" > ${target_file}
