<script setup lang="ts">
import { onMounted } from 'vue'

onMounted(() => {
  console.info('modality plugin starter is ready')
})
</script>

<template>
  <section id="modality-plugin-starter">
    <div class="wrapper">
      <span class="title"> modality 插件已启动 </span>
      <span class="message">编辑 src/starter.go 与 ui/src 开始扩展平台</span>
      <div class="guides">
        <article class="guide">
          <h2>后端服务</h2>
          <p>在 src/starter.go 中注册 gRPC 服务，并在 plugin.json 声明服务名。</p>
        </article>
        <article class="guide">
          <h2>宿主访问</h2>
          <p>只申请必要的 grants；宿主会重新鉴权并默认拒绝越权调用。</p>
        </article>
        <article class="guide">
          <h2>控制台页面</h2>
          <p>在 ui/src/index.ts 注册路由，并保持 routes.json 声明一致。</p>
        </article>
        <article class="guide">
          <h2>本地验证</h2>
          <p>运行 make verify 与 make package，再挂载到宿主插件目录调试。</p>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
#modality-plugin-starter {
  min-height: 100vh;
  background: linear-gradient(135deg, #f8fafc, #eef2ff);
  font-family: "Avenir Next", "PingFang SC", "Microsoft YaHei", sans-serif;
}

.wrapper {
  display: flex;
  min-height: 100vh;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 1.5rem;
  padding: 3rem 1.5rem;
}

.title {
  color: #172033;
  font-size: 1.5rem;
  font-weight: 700;
}

.message {
  color: #4b5563;
}

.guides {
  display: grid;
  width: min(52rem, 100%);
  grid-template-columns: repeat(1, minmax(0, 1fr));
  gap: 1rem;
}

.guide {
  border: 1px solid #dbeafe;
  border-radius: 0.75rem;
  background: rgba(255, 255, 255, 0.86);
  padding: 1rem;
  box-shadow: 0 12px 30px rgba(15, 23, 42, 0.08);
}

.guide h2 {
  color: #1d4ed8;
  font-size: 1.05rem;
  margin-bottom: 0.5rem;
}

.guide p {
  color: #475569;
  font-size: 0.9rem;
  line-height: 1.6;
}

@media (min-width: 760px) {
  .guides {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}
</style>
