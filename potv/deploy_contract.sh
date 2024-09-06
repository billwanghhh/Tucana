#!/bin/bash
transfer_value(){
  file="$HOME/.tucd/config/genesis.json"
  prefix="tuc"

  # 使用grep查找目标前缀，并使用sed提取完整字符串
  grep -o "\"${prefix}[^\"]*\"" "$file" | sed 's/"//g' > extracted_tuc_addresses.txt
  echo "Extracted addresses with prefix '$prefix' have been saved to extracted_tuc_addresses.txt."
  third_line=$(sed -n '3p' extracted_tuc_addresses.txt)
  # 输出结果
  echo "The third line is: $third_line"
  tucd tx bank send $third_line \
  tuc1pnhvqhfplttfu89rlshaphy5s4mt2jwzwefu0d 10000tuc --fees=40utuc --chain-id=tuc_9000-1 -y
  ##
  echo "wait transaction finished"
  sleep 10
  cast balance 0x0ceEC05D21fAd69e1CA3FC2fD0dc948576B549C2
}

transfer_value

rm -rf artifacts cache .openzeppelin


PRIVATE_KEY=f08a80d51fd9f9fa6cdbe450af4d2aa98a5d2632ab4847cfeac77233c911f832 \
npx hardhat run scripts/deploy.ts  --network localhost


PRIVATE_KEY=f08a80d51fd9f9fa6cdbe450af4d2aa98a5d2632ab4847cfeac77233c911f832 \
npx hardhat run scripts/configSet.ts  --network localhost


PRIVATE_KEY=f08a80d51fd9f9fa6cdbe450af4d2aa98a5d2632ab4847cfeac77233c911f832 \
npx hardhat run scripts/supply.ts  --network localhost