<?php
declare(strict_types=1);

namespace App\Command;

use App\Document\User;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Security\Core\Encoder\UserPasswordEncoderInterface;

class CreateUserCommand  extends Command
{
    /**
     * @var DocumentManager
     */
    private $documentManager;

    /**
     * @var UserPasswordEncoderInterface
     */
    private $encoder;

    /**
     * CreateUserCommand constructor.
     * @param DocumentManager $documentManager
     * @param UserPasswordEncoderInterface $encoder
     */
    public function __construct(DocumentManager $documentManager, UserPasswordEncoderInterface $encoder)
    {
        parent::__construct();
        $this->documentManager = $documentManager;
        $this->encoder = $encoder;
    }

    protected function configure()
    {
        $this->setName('app:create-user');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $user = new User();
        $user->setUsername('test2');

        $plainPassword = 'test2';
        $encoded = $this->encoder->encodePassword($user, $plainPassword);

        $user->setPassword($encoded);

        $this->documentManager->persist($user);
        $this->documentManager->flush();
    }
}
