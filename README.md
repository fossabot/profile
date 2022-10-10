# 简介

本项目的配置适用于 [**Clash Premium 内核**](https://github.com/Dreamacro/clash/releases/tag/premium)、 [**Surge**](https://nssurge.com/)、 [**Quanmutumlt X**](https://apps.apple.com/us/app/quantumult-x/id1443988620) 

## 说明

本项目中规则集的数据是每天定时拉取 [@dler-io/Rules](https://github.com/dler-io/Rules) 项目内容重新生成

#### Clash Rule Providers 配置方式

```yaml
rule-providers:
  reject:
    type: http
    behavior: domain
    url: "https://github.com/srk24/profile/raw/master/clash/provider/reject.yaml"
    path: ./ruleset/reject.yaml
    interval: 86400
```

#### Surge Domain-set 配置方式

```
DOMAIN-SET,https://github.com/srk24/profile/raw/master/surge/list/reject.list,REJECT
```

## 致谢

- [@dler-io/Rules](https://github.com/dler-io/Rules)
