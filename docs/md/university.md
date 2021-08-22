- university
  - campus
    - department
      - college
        - grade
          - major
            - class

1. 整体上看是个链表，每级都有一个头指针
2. 除`class`外每个头指针节点都有Rank和Children，记录级别和子节点指针
3. 节点属性有：uid
