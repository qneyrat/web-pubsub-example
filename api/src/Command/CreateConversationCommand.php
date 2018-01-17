<?php
declare(strict_types=1);

namespace App\Command;

use App\Entity\Conversation;
use App\Entity\Message;
use App\Entity\User;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

class CreateConversationCommand  extends Command
{
    /**
     * @var EntityManagerInterface
     */
    private $entityManager;

    /**
     * CreateConversationCommand constructor.
     * @param EntityManagerInterface $entityManager
     */
    public function __construct(EntityManagerInterface $entityManager)
    {
        parent::__construct();
        $this->entityManager = $entityManager;
    }

    protected function configure()
    {
        $this->setName('app:create-conversation');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $user = $this->entityManager->getRepository(User::class)->find(1);
        if (!$user instanceof User) {
            return;
        }

        $conversation = new Conversation();
        $conversation->addUser($user);

        $message = new Message();
        $message->setFrom($user);
        $message->setBody('hello world!');
        $message->setConversation($conversation);

        $conversation->addMessage($message);

        $this->entityManager->persist($conversation);
        $this->entityManager->flush();
    }
}
