---
layout: page
title: Index
tagline:
---

<ul class="posts">
  {% for post in site.posts %}
    <li><span>{{ post.date | date_to_string }}</span> &raquo; <a href="{{ BASE_PATH }}{{ post.url }}">{{ post.title }}</a></li>
  {% endfor %}
</ul>

>如果每个人都是一颗星球，我感激我们的光锥曾彼此重叠，而你永远改变了我的星轨。  
你是我所在的星系未曾分崩离析的原因，是我宇宙之网的永恒组成。  