# Shortify

1. Dependency Injection

   No contexto do Go, implementei o Dependency Injection de forma centralizada
   no adapter, onde todas as dependências do projeto são inseridas. Isso facilita
   a testabilidade e permite que qualquer alteração em um serviço seja refletida
   de forma consistente em toda a aplicação.

```
  	server.NewServices(
		server.WithUrlRepository(urlRepository),
		server.WithUrlCache(urlCache),
		server.WithUrlUseCase(urlUseCase),
		server.WithUrlPersistMediator(urlMediator),
	)
```

2. Factory

   Implementei o padrão Factory para facilitar a expansão e modificação das 
   funcionalidades. O objetivo é que futuramente seja possível adicionar diferentes
   tipos de planos para os usuários, cada um com regras específicas de expiração das URLs.
   Esse design respeita o princípio Open/Closed (aberto para extensões, fechado para modificações),
   permitindo que novos tipos de encurtamento de URL sejam adicionados sem a necessidade
   de alterar o código existente.

```
   func EncriptByHashMethod(method hash.Method) EncodeUrlFactory {
       return map[hash.Method]entityBase62{
           hash.Base62: NewBase62(),
       }[method]
   }
   
   func GetUserPlan(method role.User) PlanFactory {
       return map[role.User]plan{
           role.BasicUser: newPlan(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())),
       }[method]
   }

```
   Exemplo de uso:
```
	plan := factory.GetUserPlan(role.BasicUser)
	if err := u.urlCache.GetConnection().
		Set(
			ctx,
			urlRecord.GetShortenedUrl().GetValue(),
			urlRecord.GetOriginalUrl().GetValue(),
			plan.TimeToExpired()).
		Err(); err != nil {
		tlog.Warn("NewUrlRecord.Create", "Create Url in Redis")
	}
```

3. Arquitetura Baseada em Eventos e Padrão Mediator

   Utilizo uma arquitetura baseada em eventos, onde o padrão Mediator é responsável
   por orquestrar a persistência das URLs. Essa abordagem permite que as funcionalidades
   sejam desacopladas e promovam uma melhor organização e fluxo de dados no sistema.

```
  func (p *PersistUrlMediator) Notify(T any) error {
	persistUrlRecordContext := context.Background()
	if err := p.urlCache.Create(persistUrlRecordContext, T.(*entity.UrlRecord)); err != nil {
		return err
	}
	_, err := p.urRepository.InsertUrlRecord(persistUrlRecordContext, T.(*entity.UrlRecord))
	if err != nil {
		return err
	}
	return nil
}
```
4. Organização de Configurações

Todas as variáveis de ambiente estão centralizadas em um arquivo config.ENVIRONMENT.yaml,
facilitando o gerenciamento e a modificação dos parâmetros de configuração do sistema.

```
   config:
     env: development
     port: 8080
     hostname: http://shortify
   pgsql:
     host: postgres://postgres:frajolinha202@localhost:5432/shortify
     entryPoint: ./init.sql
   redis:
     url: redis://localhost:6379

```


5. Domain-Driven Design (DDD)
   O projeto segue os princípios de Domain-Driven Design (DDD), organizando o código
   de forma a refletir claramente os domínios e subdomínios da aplicação.
   Isso garante uma separação de responsabilidades bem definida e facilita a
   evolução do projeto conforme novas funcionalidades são adicionadas.

6. Demais Observacões

* swager configurado
* subdomínio de suporte `github.com/ribeirosaimon/tooltip v0.0.1`
* swager
* lint
* cobertura de testes unitários e de integracão com teste containers
   ![img.png](img.png)