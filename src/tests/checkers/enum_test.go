package checkers_test

import (
	"testing"

	"github.com/VKCOM/noverify/src/linttest"
)

func TestEnumType(t *testing.T) {
	// See #533.
	linttest.SimpleNegativeTest(t, `<?php
enum PushTemplateStatus: string
{
    case DRAFT = 'draft';
    case PLANNED = 'planned';
    case IN_PROGRESS = 'in_progress';
    case SENT = 'sent';

    public static function values(): array
    {
        return array_map(fn(self $case) => $case->value, self::cases());
    }



    public function getTranslateStatus(): string
    {
        /** @noverify-suppress undefinedConstant  */
        return match($this) {
            self::DRAFT => 'Черновик',
            self::PLANNED => 'Запланированно',
            self::IN_PROGRESS => 'В процессе',
            self::SENT => 'Отправленно',
            self::TEST => 'Отправленно',
        };
    }
}

`)
}

func TestEnumAsTypeHint(t *testing.T) {
	test := linttest.NewSuite(t)
	test.AddFile(`<?php
enum A: int {
	case ONE = 3;
}
enum B: string {
	case TWO = 'test';
}

class Foo {
  private A $a = A::ONE;
  public static A $a1;

  public function f(A $a): A {}
  public function f1(A $a, B $b): A {}
}

function f(A $a): A {}
function f1(A $a, B $b): A {
  $_ = function(A $a): B {};
}

`)
	test.Expect = []string{
		`Cannot use trait A as a typehint for property type`,
		`Cannot use trait A as a typehint for property type`,
		`Cannot use trait A as a typehint for return type`,
		`Cannot use trait A as a typehint for parameter type`,
		`Cannot use trait A as a typehint for return type`,
		`Cannot use trait A as a typehint for parameter type`,
		`Cannot use trait B as a typehint for parameter type`,
		`Cannot use trait A as a typehint for return type`,
		`Cannot use trait A as a typehint for parameter type`,
		`Cannot use trait A as a typehint for return type`,
		`Cannot use trait A as a typehint for parameter type`,
		`Cannot use trait B as a typehint for parameter type`,
		`Cannot use trait B as a typehint for closure return type`,
		`Cannot use trait A as a typehint for parameter type`,
	}
	linttest.RunFilterMatch(test, "badTraitUse")
}
